# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make differential_mutual_information_graph.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

differential_mutual_information_graph.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > differential_mutual_information_graph.tex

differential_mutual_information_graph.pdf: ../../*.tex ../../trademark.pdf *.tex differential_mutual_information_graph.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode differential_mutual_information_graph.tex
	../../latexrun differential_mutual_information_graph.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: differential_mutual_information_graph.pdf
	open differential_mutual_information_graph.pdf

o:
	make open

terms: differential_mutual_information_graph.tex
	bbk_sheet -mode ts differential_mutual_information_graph.tex

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
	rm -f differential_mutual_information_graph.*
	make reset
	bbk_sheet -mode mk > Makefile