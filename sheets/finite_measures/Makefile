# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make finite_measures.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

finite_measures.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > finite_measures.tex

finite_measures.pdf: ../../*.tex ../../trademark.pdf *.tex finite_measures.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode finite_measures.tex
	../../latexrun finite_measures.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: finite_measures.pdf
	open finite_measures.pdf

o:
	make open

terms: finite_measures.tex
	bbk sheet -mode ts finite_measures.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f finite_measures.*
	make reset
	bbk sheet -mode mk > Makefile