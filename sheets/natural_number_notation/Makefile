# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make natural_number_notation.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

natural_number_notation.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > natural_number_notation.tex

natural_number_notation.pdf: ../../*.tex ../../trademark.pdf *.tex natural_number_notation.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode natural_number_notation.tex
	../../latexrun natural_number_notation.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_number_notation.pdf
	open natural_number_notation.pdf

o:
	make open

terms: natural_number_notation.tex
	bbk_sheet -mode ts natural_number_notation.tex

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
	rm -f natural_number_notation.*
	make reset
	bbk_sheet -mode mk > Makefile