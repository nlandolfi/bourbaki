# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make number_of_elements.pdf

number_of_elements.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > number_of_elements.tex

number_of_elements.pdf: ../../*.tex ../../trademark.pdf *.tex number_of_elements.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode number_of_elements.tex
	../../latexrun number_of_elements.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: number_of_elements.pdf
	open number_of_elements.pdf

o:
	make open

terms: number_of_elements.tex
	bbk_sheet -mode ts number_of_elements.tex

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
	rm -f number_of_elements.*
	make reset
	bbk_sheet -mode mk > Makefile