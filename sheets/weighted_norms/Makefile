# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make weighted_norms.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

weighted_norms.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > weighted_norms.tex

weighted_norms.pdf: ../../*.tex ../../trademark.pdf *.tex weighted_norms.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode weighted_norms.tex
	../../latexrun weighted_norms.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: weighted_norms.pdf
	open weighted_norms.pdf

o:
	make open

terms: weighted_norms.tex
	bbk sheet -mode ts weighted_norms.tex

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
	rm -f weighted_norms.*
	make reset
	bbk sheet -mode mk > Makefile