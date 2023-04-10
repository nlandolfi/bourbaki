# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make inverse_elements.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

inverse_elements.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > inverse_elements.tex

inverse_elements.pdf: ../../*.tex ../../trademark.pdf *.tex inverse_elements.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode inverse_elements.tex
	../../latexrun inverse_elements.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: inverse_elements.pdf
	open inverse_elements.pdf

o:
	make open

terms: inverse_elements.tex
	bbk sheet -mode ts inverse_elements.tex

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
	rm -f inverse_elements.*
	make reset
	bbk sheet -mode mk > Makefile