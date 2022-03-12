# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make digital_naturals.pdf

digital_naturals.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > digital_naturals.tex

digital_naturals.pdf: ../../*.tex ../../trademark.pdf *.tex digital_naturals.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode digital_naturals.tex
	../../latexrun digital_naturals.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: digital_naturals.pdf
	open digital_naturals.pdf

o:
	make open

terms: digital_naturals.tex
	bbk_sheet -mode ts digital_naturals.tex

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
	rm -f digital_naturals.*
	make reset
	bbk_sheet -mode mk > Makefile