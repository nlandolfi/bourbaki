# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make floating_point_representations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

floating_point_representations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > floating_point_representations.tex

floating_point_representations.pdf: ../../*.tex ../../trademark.pdf *.tex floating_point_representations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode floating_point_representations.tex
	../../latexrun floating_point_representations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: floating_point_representations.pdf
	open floating_point_representations.pdf

o:
	make open

terms: floating_point_representations.tex
	bbk_sheet -mode ts floating_point_representations.tex

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
	rm -f floating_point_representations.*
	make reset
	bbk_sheet -mode mk > Makefile