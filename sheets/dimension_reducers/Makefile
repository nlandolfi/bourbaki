# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make dimension_reducers.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

dimension_reducers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > dimension_reducers.tex

dimension_reducers.pdf: ../../*.tex ../../trademark.pdf *.tex dimension_reducers.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode dimension_reducers.tex
	../../latexrun dimension_reducers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: dimension_reducers.pdf
	open dimension_reducers.pdf

o:
	make open

terms: dimension_reducers.tex
	bbk_sheet -mode ts dimension_reducers.tex

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
	rm -f dimension_reducers.*
	make reset
	bbk_sheet -mode mk > Makefile