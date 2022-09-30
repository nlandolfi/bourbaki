# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make length_common_notions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

length_common_notions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > length_common_notions.tex

length_common_notions.pdf: ../../*.tex ../../trademark.pdf *.tex length_common_notions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode length_common_notions.tex
	../../latexrun length_common_notions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: length_common_notions.pdf
	open length_common_notions.pdf

o:
	make open

terms: length_common_notions.tex
	bbk_sheet -mode ts length_common_notions.tex

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
	rm -f length_common_notions.*
	make reset
	bbk_sheet -mode mk > Makefile