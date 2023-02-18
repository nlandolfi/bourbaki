# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make neural_networks.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

neural_networks.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > neural_networks.tex

neural_networks.pdf: ../../*.tex ../../trademark.pdf *.tex neural_networks.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode neural_networks.tex
	../../latexrun neural_networks.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: neural_networks.pdf
	open neural_networks.pdf

o:
	make open

terms: neural_networks.tex
	bbk_sheet -mode ts neural_networks.tex

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
	rm -f neural_networks.*
	make reset
	bbk_sheet -mode mk > Makefile