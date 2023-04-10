# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make finite_set_examples.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

finite_set_examples.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > finite_set_examples.tex

finite_set_examples.pdf: ../../*.tex ../../trademark.pdf *.tex finite_set_examples.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode finite_set_examples.tex
	../../latexrun finite_set_examples.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: finite_set_examples.pdf
	open finite_set_examples.pdf

o:
	make open

terms: finite_set_examples.tex
	bbk sheet -mode ts finite_set_examples.tex

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
	rm -f finite_set_examples.*
	make reset
	bbk sheet -mode mk > Makefile