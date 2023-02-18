# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make multisets.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

multisets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > multisets.tex

multisets.pdf: ../../*.tex ../../trademark.pdf *.tex multisets.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode multisets.tex
	../../latexrun multisets.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: multisets.pdf
	open multisets.pdf

o:
	make open

terms: multisets.tex
	bbk_sheet -mode ts multisets.tex

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
	rm -f multisets.*
	make reset
	bbk_sheet -mode mk > Makefile